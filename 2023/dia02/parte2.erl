

-module(parte2).
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

parse_input(_, eof, Games) -> Games;

parse_input(File, {ok, Line}, Games) ->
    [_, Id | Cubes ] = string:tokens(string:trim(Line), ":;, "),
    {Red, Green, Blue} = cubes(Cubes),

    Game = #game{
              id = int(Id),
              red = Red,
              green = Green,
              blue = Blue
            },
    parse_input(File, file:read_line(File), [Game | Games]).

parse_input() ->
    {ok, File} = file:open("bigboy.txt", [read]),
    parse_input(File, file:read_line(File), []).

start() ->
    Start = erlang:timestamp(),
    Input = parse_input(),
    io:format("input: ~f~n", [timer:now_diff(erlang:timestamp(), Start) / 1000000]),
    EricWTF = [ Red * Green * Blue || 
                #game{red = Red, green = Green, blue = Blue} <- Input ],
    erlang:display(lists:sum(EricWTF)),
    io:format("time: ~f~n", [timer:now_diff(erlang:timestamp(), Start) / 1000000]).
