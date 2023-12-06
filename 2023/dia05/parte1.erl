-module(parte1).
-export([start/0]).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    [_Seeds | Sections] = 
    string:split( string:trim( binary:bin_to_list(Binary)), "\n\n", all),
    Maps =
    begin
        A = [ begin [_ | B] = string:split(X, ":\n"), B end || X <- Sections],
        B = [ string:split(X, "\n", all) || X <- A],
        C = [ lists:map(fun(Y) -> 
                                _A = string:tokens(Y, " ") ,
                                [ begin {Z, _} = string:to_integer(_X), Z end || _X <- _A]
                        end, X) || X <- B],
        C
    end,
    Seeds = 
    begin
        [_ | [A1]] = string:split(_Seeds, ": "),
        B1 = string:tokens(A1, " "),
        [ begin {Y, _} = string:to_integer(X), Y end || X <- B1]
    end,
    {Seeds, Maps}.


loop(Seed, []) -> Seed;
loop(Seed, [Map | Maps]) ->
    % erlang:display(Seed),
    Next = 
    [ Dest + (Seed - Source)
     || [Dest, Source, Range] <- Map, (Seed >= Source) and (Seed < Source + Range)],
    case Next of
        [] -> loop(Seed, Maps);
        [A] -> loop(A, Maps)
    end.

start() ->
    {Seeds, Maps } = parse_input(),
    erlang:display(lists:min([ loop(X, Maps) || X <- Seeds])).

