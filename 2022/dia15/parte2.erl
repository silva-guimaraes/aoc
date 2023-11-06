-module(parte2).
-export([start/0]).

% nao funciona mas fica aqui a tentativa

-record(sensor, 
        {
         x, y,
         bx, by,
         beaconDistance
        }).

parse_input() ->
    {ok, Binary} = file:read_file("./sensors.txt"),
    {ok, Regex} = re:compile("-?\\d+"),
    File = string:trim(binary:bin_to_list(Binary)),
    Lines = string:split(File, "\n", all),
    [
     begin 
         {match, [[A],[B],[C],[D]]} = re:run(X, Regex, [global, {capture, all, list}]),
         {SensorX, _} = string:to_integer(A),
         {SensorY, _} = string:to_integer(B),
         {BeaconX, _} = string:to_integer(C),
         {BeaconY, _} = string:to_integer(D),
         #sensor{x  = SensorX, y = SensorY, bx = BeaconX, by = BeaconY, 
                beaconDistance = 
                erlang:abs(SensorY - BeaconY) + erlang:abs(SensorX - BeaconX)}
     end
     || X <- Lines].

edges([#sensor{x=SX, y=SY , beaconDistance=Distance} = H|T]) ->
    Edge = Distance+1,
    [
     [ {SX+A, SY-Edge+A}, {SX+Edge-A, SY+A}, 
       {SX-A, SY+Edge-A}, {SX-Edge+A, SY-A} ]
     || A <- lists:seq(0, Edge) 
    ].

outside_sensors({X, Y}, Sensors) ->
    lists:all(
      fun(#sensor{x=SX, y=SY , beaconDistance=Distance}) ->
              erlang:abs(SY - Y) + erlang:abs(SX - X) > Distance
      end,
      Sensors
     ).

loop([{X, Y}|T], Sensors, Max) when (X < 0) or (X > Max) or (Y < 0) or (Y > Max) ->
    loop(T, Sensors, Max);

loop([H|T], Sensors, Max) ->
    case outside_sensors(H, Sensors) of
        true -> H;
        _ -> loop(T, Sensors, Max)
    end.

% inside({X, Y}, Max) ->
%     (X >= 0) and (X =< Max) and (Y >= 0) and (Y =< Max).


start() ->
    Max = 20,
    Sensors = parse_input(),
    lists:map(fun edges/1, Sensors).
    % Edges = lists:uniq(lists:flatmap(fun edges/1, Sensors)), % lista de listas
    % Edges.
    % loop(Edges, Sensors, Max).
    % [{X, Y}] = [ E || E <- Edges, inside(E, Max), outside_sensors(E, Sensors) ],
    % erlang:display(X * 4000000 + Y).
