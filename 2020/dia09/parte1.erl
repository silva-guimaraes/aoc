-module(parte1).
-export([
        start/0
        ]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    [begin {Y, _} = string:to_integer(X), Y end || X <- Lines].


foobar(_, []) -> undefined;
foobar([_|Tail] = Preamble, [Current|Head]) ->
    case [X + Y || X <- Preamble, Y <- Preamble, X + Y == Current] of
        [] -> Current;
        _ ->
            foobar(Tail ++ [Current], Head)
    end.

foobar(List) ->
    Tail = lists:sublist(List, 25),
    Head = lists:sublist(List, 25+1, length(List)-25+1),
    foobar(Tail, Head).


start() ->
    Input = parse_input(),
    erlang:display(foobar(Input)).



