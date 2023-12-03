-module(parte1).
-export([start/0]).
-record(game, 
        {
         id :: integer(), 
         red :: integer(), 
         green :: integer(), 
         blue :: integer() 
         }).

int(A) -> {N, _} = string:to_integer(A), N.

cubes([], Red, Green, Blue)                 -> {Red, Green, Blue};
cubes([N, "red" | T], Red, Green, Blue)     -> cubes(T, max(Red, int(N)), Green, Blue);
cubes([N, "green" | T], Red, Green, Blue)   -> cubes(T, Red, max(Green, int(N)), Blue);
cubes([N, "blue" | T], Red, Green, Blue)    -> cubes(T, Red, Green, max(Blue, int(N))).

cubes(L) -> cubes(L, 0, 0, 0).

parse_input(_, eof) -> [];

parse_input(File, {ok, Line}) ->
    % regex omegalul
    [_, Id | Cubes ] = string:tokens(string:trim(Line), ":;, "),
    {Red, Green, Blue} = cubes(Cubes),

    Game = #game{
              id = int(Id),
              red = Red,
              green = Green,
              blue = Blue
            },
    [ Game | parse_input(File, file:read_line(File)) ].

parse_input() ->
    {ok, File} = file:open("input.txt", [read]),
    parse_input(File, file:read_line(File)).

start() ->
    Input = parse_input(),
    Possible = [ Id || #game{id = Id, red = Red, green = Green, blue = Blue} <- Input,
                       Red =< 12,
                       Green =< 13,
                       Blue =< 14
               ],
    erlang:display(lists:sum(Possible)).
