-module(parte2).
-export([
        start/0
        ]).

parse(Line) ->
    [Range, [Letter|_], Password] = string:tokens(Line, " "),
    [FirstString, SecondString] = string:tokens(Range, "-"),
    {First, _} = string:to_integer(FirstString),
    {Second, _} = string:to_integer(SecondString),
    {First, Second, Letter, Password}.

valid({First, Second, Letter, Password}) ->
    FirstLetter = lists:nth(First, Password),
    SecondLetter = lists:nth(Second, Password),
    not (FirstLetter == SecondLetter) and
    ((FirstLetter == Letter) or (SecondLetter == Letter)).

start() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = binary:bin_to_list(Binary),
    Lines = string:tokens(File, "\n"),
    Passwords = [parse(X) || X <- Lines],
    Valid = length([X || X <- Passwords, valid(X)]),
    io:format("~p~n", [Valid]).

