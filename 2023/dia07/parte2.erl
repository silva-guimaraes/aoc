-module(parte2).
-export([start/0]).

translate($A) -> 14;
translate($K) -> 13;
translate($Q) -> 12;
translate($J) -> 1;
translate($T) -> 10;
translate(X) -> X - $0.

-record(hand, {hand, top, bid}).

top([], Ret) ->
    lists:reverse(lists:keysort(2, maps:to_list(Ret)));

top([H|T], Ret) ->
    top(T, Ret#{H => maps:get(H, Ret, 0) + 1}).

top(L) -> top(L, #{}).

parse_input(_, eof, Ret) -> Ret;

parse_input(File, {ok, Line}, Ret) ->
    [_Hand | _Bid] = string:split(string:trim(Line), " "),
    {Bid, _} = string:to_integer(_Bid),
    Hand = [ translate(X) || X <- _Hand ],
    Top = top(Hand),
    parse_input(File, file:read_line(File), [ #hand{hand=Hand, top=Top, bid=Bid} | Ret]).

parse_input() ->
    {ok, File} = file:open("bigboy.txt", [read]),
    parse_input(File, file:read_line(File), []).

type(#hand{top=Top}) ->
    J = 1,
    case lists:keytake(J, 1, Top) of
        false -> type(Top);
        % lista vazia significa que aviam apenas J, um five of a kind
        {value, _, []} -> 7;
        {value, {_, Jamount}, [{First, TopAmount} | _] = New} -> 
            type(lists:keystore(First, 1, New, {First, TopAmount + Jamount}))
    end;

type([_])                   -> 7;
type([{_, 3}, _])           -> 5;
type([_, _])                -> 6;
type([{_, A}, {_, A}, _])   -> 3;
type([_, _, _])             -> 4;
type([{_, 2} | _])          -> 2;
type(_)                     -> 1.

winning([], _) -> 0;
winning([#hand{bid=Bid} | T], Rank) ->
    Bid * Rank + winning(T, Rank+1).

compare(#hand{hand=HandA}, #hand{hand=HandB}) ->
    compare(HandA, HandB);

compare([A | _], [B | _]) when A > B -> false;
compare([A | _], [B | _]) when B > A -> true;
compare([A | AT], [A | BT]) -> 
    compare(AT, BT).

loop([], _) -> 0;
loop([{_, H}|T], Rank) ->
    Sorted =
    lists:sort(fun compare/2, H),
    winning(Sorted, Rank) + loop(T, Rank + length(H)).

loop(L) ->
    loop(L, 1).


start() ->
    Input = parse_input(),
    Start = erlang:timestamp(),
    Grouped = maps:to_list(maps:groups_from_list(fun type/1, Input)),
    erlang:display(loop(Grouped)),
    io:format("~f~n", [timer:now_diff(erlang:timestamp(), Start) / 1000000]).
