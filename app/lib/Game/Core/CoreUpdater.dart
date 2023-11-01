import 'package:orbstrike/Game/Player/Player.dart';
import 'package:orbstrike/Game/World/WorldBackground.dart';
import 'package:orbstrike/Game/World/WorldBorder.dart';
import 'package:orbstrike/Game/Lib/Types/GameCoreComponents.dart';

import 'package:flame/components.dart';
import 'package:flutter/material.dart';

/// Changes the state of the playerComponents and the mainPlayerComponent based on the new GameBoard
///
/// Returns a Map containing Components to add (true) and remove (false) from the world
Map<Component, bool> updateGameBoard(final GameCoreComponents coreComp) {
  Map<Component, bool> componentBuffer = {};

  // Create main component if not existent
  if (coreComp.mainPlayerComponent==null && coreComp.board.players[coreComp.mainPlayerCreds?.id]!=null) {
    coreComp.mainPlayerComponent
    = MainPlayerComponent(
      networkPlayerRep: coreComp.board.players[coreComp.mainPlayerCreds?.id]!,
      collided: coreComp.mainPlayerCollided,
      lerpFactor: coreComp.lerpFactor,
      playerNormalSS: coreComp.playerNameSS,
      playerRingedSS: coreComp.playerRingedSS,
      ringAnimateSS: coreComp.ringAnimateSS,
    );

    coreComp.world.add(coreComp.mainPlayerComponent!);
  }

  // Update main component
  final player = coreComp.board.players[coreComp.mainPlayerCreds?.id];
  if (player!=null && coreComp.mainPlayerComponent!=null) {
    coreComp.mainPlayerComponent?.networkPlayerRep = player;
  }

  // Remove players that are not existent anymore
  coreComp.playerComponents.removeWhere((id, component) {
    if (!coreComp.board.players.containsKey(id)) {
      componentBuffer.putIfAbsent(component, () => false);
      return true;
    }
    return false;
  });

  // Add players that were added and update the existing
  for (var networkPlayerRep in coreComp.board.players.values) {
    if (!coreComp.playerComponents.containsKey(networkPlayerRep.id)) {
      final player = EnemyPlayerComponent(
        networkPlayerRep: networkPlayerRep,
        lerpFactor: coreComp.lerpFactor,
        playerNormalSS: coreComp.playerNameSS,
        playerRingedSS: coreComp.playerRingedSS,
        ringAnimateSS: coreComp.ringAnimateSS,
      );
      if (networkPlayerRep.id!=coreComp.mainPlayerCreds?.id) {
        coreComp.playerComponents[networkPlayerRep.id] = player;
        componentBuffer.putIfAbsent(player, () => true);
      }
    } else {
      if (networkPlayerRep.id!=coreComp.mainPlayerCreds?.id) {
        coreComp.playerComponents[networkPlayerRep.id]?.networkPlayerRep = networkPlayerRep;
      }
    }
  }

  // Update Game border if necessary
  if (coreComp.border==null) {
    coreComp.border = WorldBorder(colors: [Colors.orange, Colors.red], radius: coreComp.board.rad, stroke: 10);
    coreComp.world.add(coreComp.border!);
  }

  // Update Game background if necessary
  if (coreComp.background==null) {
    coreComp.background = WorldBackground(
      radius: coreComp.board.rad,
      rectSize: 5,
      rectBorderRadius: 12,
      rectPaint: Paint()..color=const Color.fromRGBO(54,57,62, 0.7),
      rectSpacing: 50,
    );
    coreComp.world.add(coreComp.background!);
  }

  return componentBuffer;
}