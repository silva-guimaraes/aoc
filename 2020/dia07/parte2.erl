-module(parte2).
-export([
        start/0
        ]).

filter_bag_names("bags")        -> false;
filter_bag_names("bag")         -> false;
filter_bag_names("contain")     -> false;
filter_bag_names("no")          -> false;
filter_bag_names("other")       -> false;
filter_bag_names(_)             -> true.


join_bag_names([])       -> [];
join_bag_names([Number, A, B|T]) ->
    {N, _} = string:to_integer(Number),
    [{N, A ++ " " ++ B} | join_bag_names(T)].

bag([First, Second|T] = _) ->
    {First ++ " " ++ Second, join_bag_names(T)}.

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    Tokens = [string:tokens(X, ", .") || X <- Lines],
    Filtered = [lists:filter(fun filter_bag_names/1, X) || X <- Tokens],
    [bag(X) || X <- Filtered ].
    % [join_bag_names(X) || X <- Bags].


build_graph([]) -> #{};
build_graph([{From, To}|T]) -> 
    maps:merge(#{From => To}, build_graph(T)).


traverse_graph(Current, Graph) ->
    case Graph of
        #{Current := []} -> 1;
        #{Current := Vertices} ->
            Sum = [traverse_graph(Bag, Graph) * Amount || {Amount, Bag} <- Vertices],
            erlang:display({Current, Sum, 1}),
            lists:sum(Sum) + 1
    end.

traverse_graph(Graph) ->
    traverse_graph("shiny gold" , Graph) - 1. %% off by one!!!
%% não sei onde esta o erro


start() ->
    Bags = parse_input(),
    Graph = build_graph(Bags),
    traverse_graph(Graph).



