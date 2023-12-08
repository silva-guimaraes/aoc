-module(parte2).
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

find_z(_, _, _, [_, _, $Z], Count) when Count > 0 -> Count;
find_z([], LeftRight, Map, Current, Count) ->
    find_z(LeftRight, LeftRight, Map, Current, Count);
find_z([H|T], LeftRight, Map, Current, Count) ->
    case H of 
        $R ->
            #{Current := [_, Next]} = Map;
        $L ->
            #{Current := [Next, _]} = Map
    end,
    find_z(T, LeftRight, Map, Next, Count+1).

find_z(L, Start, Map) ->
    find_z(L, L, Map, Start, 0).

lcm(X, T, I) when (X*I) rem T == 0 -> X*I;
lcm(X, T, I) -> lcm(X, T, I+1).
lcm(X, T) -> lcm(X, T, 1).

start() ->
    {LeftRight, Map} = parse_input(),
    Z = [ X || X = [_, _, $Z] <- maps:keys(Map)],
    [H|T] = [find_z(LeftRight, X, Map) || X <- Z],
    lists:foldl(fun lcm/2, H, T).
