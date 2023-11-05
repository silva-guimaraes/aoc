-module(parte1).
-export([start/0]).


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


merge_slices([A], Ret) -> [A | Ret];

% -----a--------------a-----
% ---------b----b-----------
merge_slices([{A1, A2} = A, {B1, B2} |T], Ret) when (A1 =< B1) and (A2 >= B2) ->
    erlang:display(A),
    merge_slices([A | T], Ret);

% -----a--------a----------
% ---------b--------b------
merge_slices([{A1, A2}, {B1, B2} |T], Ret) when (B1 >= A1) and (B1 =< A2) ->
    % erlang:display({A1, B2}),
    merge_slices([{A1, B2} | T], Ret).

% % --a--------a------------------
% % -----------------b--------b---
% merge_slices([H|T], Ret) ->
%     % erlang:display(H),
%     merge_slices([T], [H|Ret]).

merge_slices(Slices) ->
    merge_slices(Slices, []).


start() ->
    Y = 2000000 ,
    % Y = 10 ,
    Sensors = parse_input(),
    Slices = [
                 begin
                     % pega a distancia até o Y alvo e subtrai pela distancia
                     % sensor-beacon. isso nos da uma metade do comprimento
                     % da area que faz intercesão com o Y. multiplicar por
                     % dois para conseguir a outra metada e adionar mais um
                     % para incluir o ponto do meio.
                     Ydistance = erlang:abs(SensorY - Y),
                     Remainder = erlang:abs(Ydistance - Distance),
                     {SensorX - Remainder, SensorX + Remainder}
                 end
                 || #sensor{y = SensorY, x = SensorX, 
                            beaconDistance = Distance} <- Sensors,
                    erlang:abs(SensorY - Y) =< Distance
                ],
    Sorted = lists:usort(fun({A, _}, {B, _}) -> A < B end, Slices),
    erlang:display(Sorted),
    [{A, B}] = merge_slices(Sorted),
    erlang:display(erlang:abs(A) + B) .
