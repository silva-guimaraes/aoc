-module(parte1).
-export([start/0]).


map([], Map) -> Map;
map([[Node | Edges]|T], Map) ->
    map(T, Map#{Node => Edges}).

map(L) -> map(L, #{}).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    [LeftRight, Nodes] = string:split(string:trim( binary:bin_to_list(Binary)), "\n\n"),
    Map = map([string:tokens(X, " =(),") || X <- string:split(Nodes, "\n", all) ]),
    {LeftRight, Map}.

loop(_, _, _, "ZZZ", Count) -> Count;
loop([], LeftRight, Map, Current, Count) ->
    loop(LeftRight, LeftRight, Map, Current, Count);

loop([H|T], LeftRight, Map, Current, Count) ->
    case H of 
        $R ->
            #{Current := [_, Next]} = Map;
        $L ->
            #{Current := [Next, _]} = Map
    end,
    loop(T, LeftRight, Map, Next, Count+1).

loop(L, Map) ->
    loop(L, L, Map, "AAA", 0).


start() ->
    {LeftRight, Map} = parse_input(),
    erlang:display(loop(LeftRight, Map)).
