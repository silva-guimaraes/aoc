

-module(parte12).
-export([start/0]).


quad(B, C) ->
    Delta = math:sqrt(B*B - 4 * -1 * -C),
    High = erlang:ceil((-B - Delta) / -2),
    Low = erlang:floor((-B + Delta) / -2),
    (High - Low) - 1.

start() ->
    io:format("parte 1:~n"),
    erlang:display(quad(56717999, 334113513502430)),
    io:format("parte 2:~n"),
    erlang:display(quad(56, 334) * quad(71, 1135) * 
                   quad(79, 1350) * quad(99, 2430)).
