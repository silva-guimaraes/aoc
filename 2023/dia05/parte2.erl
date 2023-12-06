-module(parte2).
-export([start/0]).


pair([]) -> [];
pair([Start, Range | T]) ->
    [{Start, Start + Range} | pair(T)].

parse_input() ->
    {ok, Binary} = file:read_file("teste.txt"),
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
%     erlang:display("foo"),
%     [
%      {pass(Start, Line), pass(Source, Line)}, 
%      {pass(Source, Line), pass(End, Line)}
%     ];
%
% find_range({Start, End}, [[_, Source, Range] = Line | _]) 
%   when (Start >= Source) and (Start < Source + Range) and
%        (End >= Source) and (End < Source + Range) ->
%     {pass(Start, Line), pass(End, Line)};
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

find_range(Range, []) -> Range;

find_range({Start, End}, [[_, Source, Range] = Line | _]) 
  when (Source =< Start) and (Start < Source + Range) and
       (End >= Source + Range) ->
    [
     {pass(Start, Line), pass(Source + Range, Line)}, 
     {pass(Source + Range, Line), pass(End, Line)}
    ];

find_range({Start, End}, [[_, Source, Range] = Line | _]) 
  when (Start < Source) and 
       (End >= Source) and (End < Source + Range) ->
    erlang:display("foo"),
    [
     {pass(Start, Line), pass(Source, Line)}, 
     {pass(Source, Line), pass(End, Line)}
    ];

find_range({Start, End}, [[_, Source, Range] = Line | _]) 
  when (Start >= Source) and (Start < Source + Range) and
       (End >= Source) and (End < Source + Range) ->
    {pass(Start, Line), pass(End, Line)};

find_range({Start, End}, [[_, Source, Range] = Line | _]) 
  when (Start < Source) and (End > Source + Range) ->
    [
     {pass(Source, Line), pass(Source + End, Line)}, 
     {pass(Start, Line), pass(Source, Line)},
     {pass(Source + Range, Line), pass(End, Line)}
    ];

find_range(Range, [_ | Lines]) ->
    find_range(Range, Lines).

% loop(Seeds, [Map | Maps]) ->
%
%     .

start() ->
    {Seeds, Maps } = parse_input(),
    find_range(lists:nth(1, Seeds), lists:nth(1, Maps)).
    % .
