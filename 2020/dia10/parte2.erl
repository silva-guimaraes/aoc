-module(parte2).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("teste.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    [begin {Y, _} = string:to_integer(X), Y end || X <- Lines].

get_diff([_]) -> [];
get_diff([A, B|T]) ->
    [B - A| get_diff([B|T])].


% count_diff([], X) -> X;
% count_diff([H|T], X) ->
%     #{H := N} = X,
%     count_diff(T, X#{H := N+1}).

% count_diff(List) -> 
%     count_diff(List, #{1 => 0, 2 => 0, 3 => 0}).

find_sublist(List, Start, _, Foo) when Start > length(List) -> Foo;
find_sublist(List, Start, Len, Foo) ->
    Sublist = lists:sublist(List, Start, Len),
    case lists:sum(Sublist) of
        X when X == 3 -> 
            find_sublist(List, Start+1, 1, Foo+1);
        X when X == 3 -> 
            erlang:display(Sublist),
            find_sublist(List, Start+1, 1, Foo+1);
        _ ->
            erlang:display(Sublist),
            find_sublist(List, Start, Len+1, Foo+1)
    end.
find_sublist(List) -> find_sublist(List, 1, 1, 0).

start() ->
    Input = parse_input(),
    Adapters = lists:sort(Input),
    Diff = get_diff(Adapters),
    erlang:display(Diff),
    find_sublist(Diff).
