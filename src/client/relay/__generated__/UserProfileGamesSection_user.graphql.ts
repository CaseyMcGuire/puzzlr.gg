/**
 * @generated SignedSource<<849d2768b0ded96758ddbc9766d975fb>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment } from 'relay-runtime';
export type GamePlayerKind = "AI" | "HUMAN" | "%future added value";
export type GameStatus = "DRAW" | "IN_PROGRESS" | "PENDING" | "WON" | "%future added value";
export type GameType = "TIC_TAC_TOE" | "%future added value";
import { FragmentRefs } from "relay-runtime";
export type UserProfileGamesSection_user$data = {
  readonly games: ReadonlyArray<{
    readonly currentTurnPlayer: {
      readonly id: string;
      readonly kind: GamePlayerKind;
      readonly user: {
        readonly id: string;
      } | null | undefined;
    } | null | undefined;
    readonly id: string;
    readonly players: ReadonlyArray<{
      readonly id: string;
      readonly kind: GamePlayerKind;
      readonly marker: string;
      readonly user: {
        readonly email: string;
        readonly id: string;
      } | null | undefined;
    }> | null | undefined;
    readonly status: GameStatus;
    readonly type: GameType;
    readonly winnerPlayer: {
      readonly id: string;
      readonly kind: GamePlayerKind;
      readonly user: {
        readonly id: string;
      } | null | undefined;
    } | null | undefined;
  }>;
  readonly id: string;
  readonly " $fragmentType": "UserProfileGamesSection_user";
};
export type UserProfileGamesSection_user$key = {
  readonly " $data"?: UserProfileGamesSection_user$data;
  readonly " $fragmentSpreads": FragmentRefs<"UserProfileGamesSection_user">;
};

const node: ReaderFragment = (function(){
var v0 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "kind",
  "storageKey": null
},
v2 = [
  (v0/*: any*/),
  (v1/*: any*/),
  {
    "alias": null,
    "args": null,
    "concreteType": "User",
    "kind": "LinkedField",
    "name": "user",
    "plural": false,
    "selections": [
      (v0/*: any*/)
    ],
    "storageKey": null
  }
];
return {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "UserProfileGamesSection_user",
  "selections": [
    (v0/*: any*/),
    {
      "alias": null,
      "args": null,
      "concreteType": "Game",
      "kind": "LinkedField",
      "name": "games",
      "plural": true,
      "selections": [
        (v0/*: any*/),
        {
          "alias": null,
          "args": null,
          "kind": "ScalarField",
          "name": "type",
          "storageKey": null
        },
        {
          "alias": null,
          "args": null,
          "kind": "ScalarField",
          "name": "status",
          "storageKey": null
        },
        {
          "alias": null,
          "args": null,
          "concreteType": "GamePlayer",
          "kind": "LinkedField",
          "name": "winnerPlayer",
          "plural": false,
          "selections": (v2/*: any*/),
          "storageKey": null
        },
        {
          "alias": null,
          "args": null,
          "concreteType": "GamePlayer",
          "kind": "LinkedField",
          "name": "currentTurnPlayer",
          "plural": false,
          "selections": (v2/*: any*/),
          "storageKey": null
        },
        {
          "alias": null,
          "args": null,
          "concreteType": "GamePlayer",
          "kind": "LinkedField",
          "name": "players",
          "plural": true,
          "selections": [
            (v0/*: any*/),
            (v1/*: any*/),
            {
              "alias": null,
              "args": null,
              "kind": "ScalarField",
              "name": "marker",
              "storageKey": null
            },
            {
              "alias": null,
              "args": null,
              "concreteType": "User",
              "kind": "LinkedField",
              "name": "user",
              "plural": false,
              "selections": [
                (v0/*: any*/),
                {
                  "alias": null,
                  "args": null,
                  "kind": "ScalarField",
                  "name": "email",
                  "storageKey": null
                }
              ],
              "storageKey": null
            }
          ],
          "storageKey": null
        }
      ],
      "storageKey": null
    }
  ],
  "type": "User",
  "abstractKey": null
};
})();

(node as any).hash = "898bef36d8b2682fcd2957fb8866bd98";

export default node;
