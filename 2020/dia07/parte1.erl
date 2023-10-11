-module(parte1).
-export([
        start/0
        ]).

filter_bag_names("bags")        -> false;
filter_bag_names("bag")         -> false;
filter_bag_names("contain")     -> false;
filter_bag_names("no")          -> false;
filter_bag_names("other")       -> false;
filter_bag_names(Number) when length(Number) =:= 1  -> false;
filter_bag_names(_) -> true.


join_bag_names([])       -> [];
join_bag_names([A, B|T]) ->
    [A ++ " " ++ B | join_bag_names(T)].

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    %% não sabia como fazer o regex disso.
    Rules = string:split(File, "\n", all),
    Words = [string:tokens(X, ", .") || X <- Rules],
    Bags = [lists:filter(fun filter_bag_names/1, X) || X <- Words],
    [join_bag_names(X) || X <- Bags].


reverse_direction(_, []) -> #{};
reverse_direction(From, [To|Tail]) ->
    maps:merge(#{To => [From]}, reverse_direction(From, Tail)).

build_graph([]) -> #{};
build_graph([H|T]) -> 
    [From|To] = H,
    maps:merge_with(
      fun(_, A, B) -> A ++ B end,
      reverse_direction(From, To), 
      build_graph(T)).


traverse_graph(_, [], Visited) -> Visited;
traverse_graph(Graph, [H|Queue], Visited) ->
    Vertices = try #{H := X} = Graph, X catch _:_ -> [] end,
    traverse_graph(Graph, lists:uniq(Queue ++ Vertices), [H|Visited]).

traverse_graph(#{"shiny gold" := Vertices} = Graph) ->
    traverse_graph(Graph, Vertices, []).


start() ->
    Bags = parse_input(),
    Graph = build_graph(Bags),
    erlang:display(length(lists:uniq(traverse_graph(Graph)))).



