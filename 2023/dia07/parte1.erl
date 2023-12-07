-module(parte1).
-export([start/0]).

translate($A) -> 14;
translate($K) -> 13;
translate($Q) -> 12;
translate($J) -> 11;
translate($T) -> 10;
translate(X) -> X - $0.

-record(hand, {hand, top, bid}).

top([], Ret) ->
    lists:sort(
      fun
          ({_, A}, {_, B}) when A > B -> true;
          (_, _) -> false 
      end, maps:to_list(Ret));

top([H|T], Ret) ->
    top(T, Ret#{H => maps:get(H, Ret, 0) + 1}).

top(L) -> top(L, #{}).

parse_input(_, eof) -> [];

parse_input(File, {ok, Line}) ->
    [_Hand | _Bid] = string:split(string:trim(Line), " "),
    {Bid, _} = string:to_integer(_Bid),
    Hand = [ translate(X) || X <- _Hand ],
    Top = top(Hand),
    % Unique = length(Top),
    [ #hand{hand=Hand, top=Top, bid=Bid} | parse_input(File, file:read_line(File)) ].

parse_input() ->
    {ok, File} = file:open("input.txt", [read]),
    parse_input(File, file:read_line(File)).

type(#hand{top=Top}) 
  when length(Top) == 1                 -> 7;

type(#hand{top=[{_, N} | _] = Top}) 
  when (length(Top) == 2) and (N == 3)  -> 5;

type(#hand{top=Top}) 
  when length(Top) == 2                 -> 6;

type(#hand{top=[{_, A}, {_, B} | _] = Top}) 
  when (length(Top) == 3) and (A == B)  -> 3;

type(#hand{top=Top}) 
  when length(Top) == 3                 -> 4;

type(#hand{top=[{_, A} | _]}) 
  when A == 2                           -> 2;

type(_) -> 1.

winning([], _) -> 0;
winning([#hand{bid=Bid} | T], Rank) ->
    Bid * Rank + winning(T, Rank+1).

compare(#hand{hand=HandA}, #hand{hand=HandB}) ->
    Foo = compare(HandA, HandB),
    Foo;

compare([A | _], [B | _]) when A > B -> false;
compare([A | _], [B | _]) when B > A -> true;
compare([A | AT], [B | BT]) when A == B -> 
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
    Grouped = maps:to_list(maps:groups_from_list(fun type/1, Input)),
    erlang:display(loop(Grouped)).
    % io:format("hello world~n").
