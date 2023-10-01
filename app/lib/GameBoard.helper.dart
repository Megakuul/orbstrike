import 'package:flame/components.dart';
import 'package:orbstrike/proto/game.pb.dart';
import 'proto/game.pbgrpc.dart';

import 'package:orbstrike/PlayerC.dart';

/// Changes the state of the playerComponents and the mainPlayerComponent based on the new GameBoard
///
/// Returns a Map containing Components to add (true) and remove (false) from the world
Map<Component, bool> updateGameBoard(final GameBoard board, final Map<int, PlayerO> playerComps, final int playerID, PlayerC? mainPlayerComponent) {
  Map<Component, bool> componentBuffer = {};

  final player = board.players[playerID];
  if (player!=null && mainPlayerComponent!=null) {
    mainPlayerComponent.pPlayer.x = player.x;
    mainPlayerComponent.pPlayer.y = player.y;
    mainPlayerComponent.pPlayer.kills = player.kills;
    mainPlayerComponent.pPlayer.color = player.color;
    mainPlayerComponent.pPlayer.ringEnabled = player.ringEnabled;
  }

  // Remove players that are not existent anymore
  playerComps.removeWhere((id, component) {
    if (!board.players.containsKey(id)) {
      componentBuffer.putIfAbsent(component, () => false);
      return true;
    }
    return false;
  });

  for (var pPlayer in board.players.values) {
    if (!playerComps.containsKey(pPlayer.id)) {
      final player = PlayerO(pPlayer: pPlayer);
      if (pPlayer.id!=playerID) {
        playerComps[pPlayer.id] = player;
        componentBuffer.putIfAbsent(player, () => true);
      }
    }
  }

  return componentBuffer;
}