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


pass(Seed, [Dest, Source, _]) ->
    Dest + (Seed - Source).

find_range(Range, []) -> Range;

find_range({Start, End}, [[_, Source, Range] = Line | _]) 
  when (Start >= Source) and (Start < Source + Range) and
       (End >= Source) and (End < Source + Range) ->
    [{pass(Start, Line), pass(End, Line)}];

find_range(Range, [_ | Lines]) ->
    find_range(Range, Lines).

slice_range(Range, []) -> [Range];

slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Source =< Start) and (Start < Source + Range) and
       (End >= Source + Range) ->
    [
     {Start, Source + Range - 1}, 
     {Source + Range, End}
    ];

slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start < Source) and 
       (End >= Source) and (End < Source + Range) ->
    [
     {Start, Source - 1}, 
     {Source, End}
    ];

slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start >= Source) and (Start < Source + Range) and
       (End >= Source) and (End < Source + Range) ->
    [{Start, End}];

slice_range({Start, End}, [[_, Source, Range] | _]) 
  when (Start < Source) and (End > Source + Range) ->
    [
     {Source, Source + Range - 1}, 
     {Start, Source - 1},
     {Source + Range, End}
    ];

slice_range(Range, [_ | Lines]) ->
    slice_range(Range, Lines).

loop(Seeds, []) -> Seeds;
loop(Seeds, [Map | Maps] = M) ->
    A = lists:map(fun(X) -> slice_range(X, Map) end, Seeds),
    B = lists:flatten(A),
    case lists:flatten(lists:map(fun(X) -> slice_range(X, Map) end, B)) of 
        X when length(B) =/= length(X) ->
            loop(X, M);
        _ -> 
            C = lists:map(fun(X) -> find_range(X, Map) end, B),
            loop(lists:flatten(C), Maps)
    end.

start() ->
    {Seeds, Maps } = parse_input(),
    Foo = loop(Seeds, Maps),
    erlang:display(lists:min([ X || {X, _} <- Foo])).
