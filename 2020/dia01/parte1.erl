-module(parte1).
-export([
        start/0
        ]).

start() ->
    {ok, Binary} = file:read_file("input.txt"),
    List = binary:bin_to_list(Binary),
    Tokens = string:tokens(List, "\n"),
    Converted = [string:to_integer(X) || X <- Tokens],
    Input = [X || {X, _} <- Converted],
    [Ret | _] = [ X*Y || X <- Input, Y <- Input, X + Y == 2020 ],
    % io:format(Ret).
    io:format("~p~n", [Ret]).

