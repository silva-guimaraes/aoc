-module(parte2).
-export([start/0]).


pair([]) -> [];
pair([Start, Range | T]) ->
    [{Start, Start + Range} | pair(T)].

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    [_Seeds | Sections] = 
    string:split( string:trim( binary:bin_to_list(Binary)), "\n\n", all),
    Maps =
    [ begin 
          [_ | B] = string:split(X, ":\n"), 
          C = string:split(B, "\n", all),
          [
           begin
               E = string:tokens(D, " "),
               lists:map(fun(F) -> {G, _} = string:to_integer(F), G end, E)
           end
           || D <- C 
          ]
      end 
      || X <- Sections],
    Seeds = 
    begin
        [_ | [A1]] = string:split(_Seeds, ": "),
        B1 = string:tokens(A1, " "),
        C1 = [ begin {Y, _} = string:to_integer(X), Y end || X <- B1],
        pair(C1)
    end,
    {Seeds, Maps}.


% pass(Seed, [Dest, Source, _] = Line) when Source > Seed ->
%     erlang:display(Seed),
%     erlang:display(Line);
pass(Seed, [Dest, Source, _]) ->
    Dest + (Seed - Source).

% find_range(Range, []) -> Range;
%
% find_range({Start, End}, [[_, Source, Range] = Line | _]) 
%   when (Source =< Start) and (Start < Source + Range) and
%        (End >= Source + Range) ->
%     [
%      {pass(Start, Line), pass(Source + Range, Line)}, 
%      {pass(Source + Range, Line), pass(End, Line)}
%     ];
%
% find_range({Start, End}, [[_, Source, Range] = Line | _]) 
%   when (Start < Source) and 
%        (End >= Source) and (End < Source + Range) ->
%     % erlang:display("foo"),
%     [
%      {pass(Start, Line), pass(Source, Line)}, 
%      {pass(Source, Line), pass(End, Line)}
%     ];
%
% find_range({Start, End}, [[_, Source, Range] = Line | _]) 
%   when (Start >= Source) and (Start < Source + Range) and
%        (End >= Source) and (End < Source + Range) ->
%     [{pass(Start, Line), pass(End, Line)}];
%
% find_range({Start, End}, [[_, Source, Range] = Line | _]) 
%   when (Start < Source) and (End > Source + Range) ->
%     [
%      {pass(Source, Line), pass(Source + End, Line)}, 
%      {pass(Start, Line), pass(Source, Line)},
%      {pass(Source + Range, Line), pass(End, Line)}
%     ];
%
% find_range(Range, [_ | Lines]) ->
%     find_range(Range, Lines).

split_range(Range, []) -> [Range];

split_range({Start, End}, [[_, Source, Range] | _]) 
  when (Source =< Start) and (Start < Source + Range) and
       (End >= Source + Range) ->
    [
     {Start, Source + Range - 1}, 
     {Source + Range, End}
    ];

split_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start < Source) and 
       (End >= Source) and (End < Source + Range) ->
    % erlang:display("foo"),
    [
     {Start, Source - 1}, 
     {Source, End}
    ];

split_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start >= Source) and (Start < Source + Range) and
       (End >= Source) and (End < Source + Range) ->
    [{Start, End}];

split_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start < Source) and (End > Source + Range) ->
    [
     {Source, Source + End - 1}, 
     {Start, Source - 1},
     {Source + Range, End}
    ];

split_range(Range, [_ | Lines]) ->
    split_range(Range, Lines).

loop(Seeds, []) -> Seeds;
loop(Seeds, [Map | Maps]) ->
    % erlang:display(Seeds),
    A = lists:map(fun(X) -> split_range(X, Map) end, Seeds),
    B = lists:flatten(A),
    C = lists:map(fun(X) -> find_range(X, Map) end, B),
    loop(lists:flatten(C), Maps).

start() ->
    {Seeds, Maps } = parse_input(),
    Foo = loop(Seeds, Maps),
    lists:min([ X || {X, _} <- Foo, X =/= 0]).
    % lists:min().
    % split_range(lists:nth(2, Seeds), lists:nth(1, Maps)).
    % .
