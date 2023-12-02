

-module(parte2).
-export([start/0]).

parse_input(_, eof) -> [];
parse_input(File, {ok, Line}) ->
    [
     string:trim(Line) | parse_input(File, file:read_line(File))
    ].

parse_input() ->
    {ok, File} = file:open("input.txt", [read]),
    parse_input(File, file:read_line(File)).


read_numbers([]) -> [];

read_numbers([H|T]) when H < $a  -> 
    [H | read_numbers(T)];

read_numbers([_|T] = L)  -> 
    Numbers = [{"one", $1}, {"two", $2}, {"three", $3}, {"four", $4}, {"five", $5} , 
               {"six", $6} , {"seven", $7}, {"eight", $8}, {"nine", $9}],
    % retorna [] caso nenhum desses numeros sejam encontrados
    Number = [ N || {X, N} <- Numbers,
                 case string:find(L, X) of
                     L -> true;
                     _ -> false
                 end ],
    case Number of
        [N] ->
            [N | read_numbers(T)];
        _ -> read_numbers(T)
    end.

first_last(Numbers) ->
    String = [lists:nth(1, Numbers), lists:last((Numbers))],
    {Number, _} = string:to_integer(String),
    Number.


start() ->
    Foo = lists:map(fun read_numbers/1, parse_input()),
    erlang:display(lists:sum(lists:map(fun first_last/1, Foo))).
