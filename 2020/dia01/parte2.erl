-module(parte2).
-export([
        start/0
        ]).

start() ->
    {ok, Binary} = file:read_file("input.txt"),
    List = binary:bin_to_list(Binary),
    Tokens = string:tokens(List, "\n"),
    Converted = [string:to_integer(X) || X <- Tokens],
    Input = [X || {X, _} <- Converted],
    [Ret | _] = [ X*Y*Z || X <- Input, 
                           Y <- Input, 
                           Z <- Input, 
                           X + Y + Z == 2020 ],
    io:format("~p~n", [Ret]).

