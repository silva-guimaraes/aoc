-module(parte2).
-export([start/0]).

% todas as sequencias de inicio e fim são desconectas.
% convenientemente todas as sequencias entram e um loop depois de chegar no primero Z.
% ainda mais conveniente é o fato de que o tamanho do loop e a distancia até o primeiro Z 
% são iguais.
% todas as sequencias possuem um tamanho de loop diferente.
% ainda mutio mais conveniente é que todos os loops terminam em um numero multiplo do
% tamanho das instruções, ou seja, o proximo loop fara o exato mesmo caminho e tera o
% exato mesmo tamanho sempre.
% basta usar minimo multiplo comum.
% imagine 6 e 4, agora tente imaginar quais dois numeros podem multiplicar 6 e 4 para
% que o produto dos dois seja o mesmo. 6x = 4y resumidamente.
% esse era o problema em mãos
% 
% moral dessa história: verifiquem bem os seus inputs. sempre procure por padrões.


map([], Map) -> Map;
map([[Node | Edges]|T], Map) ->
    map(T, Map#{Node => Edges}).

map(L) -> map(L, #{}).

parse_input() ->
    {ok, Binary} = file:read_file("input.txt"),
    [LeftRight, Nodes] = string:split(string:trim( binary:bin_to_list(Binary)), "\n\n"),
    Map = map([string:tokens(X, " =(),") || X <- string:split(Nodes, "\n", all) ]),
    {LeftRight, Map}.

find_z(_, _, _, [_, _, $Z], Count) when Count > 0 -> Count;
find_z([], LeftRight, Map, Current, Count) ->
    find_z(LeftRight, LeftRight, Map, Current, Count);
find_z([H|T], LeftRight, Map, Current, Count) ->
    case H of 
        $R ->
            #{Current := [_, Next]} = Map;
        $L ->
            #{Current := [Next, _]} = Map
    end,
    find_z(T, LeftRight, Map, Next, Count+1).

find_z(L, Start, Map) ->
    find_z(L, L, Map, Start, 0).

% mesmo tendo dito aquilo essa não é a melhor imprementação de mmc...
lcm(X, T, I) when (X*I) rem T == 0 -> X*I;
lcm(X, T, I) -> lcm(X, T, I+1).
lcm(X, T) -> lcm(X, T, 1).

start() ->
    {LeftRight, Map} = parse_input(),
    Z = [ X || X = [_, _, $Z] <- maps:keys(Map)],
    [H|T] = [find_z(LeftRight, X, Map) || X <- Z],
    lists:foldl(fun lcm/2, H, T).
