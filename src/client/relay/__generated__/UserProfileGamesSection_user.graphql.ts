/**
 * @generated SignedSource<<a1f76c2757e5c0ae1ba93221987187d8>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment } from 'relay-runtime';
export type GameStatus = "DRAW" | "IN_PROGRESS" | "PENDING" | "WON" | "%future added value";
export type GameType = "TIC_TAC_TOE" | "%future added value";
import { FragmentRefs } from "relay-runtime";
export type UserProfileGamesSection_user$data = {
  readonly games: ReadonlyArray<{
    readonly currentTurn: {
      readonly id: string;
    } | null | undefined;
    readonly id: string;
    readonly status: GameStatus;
    readonly type: GameType;
    readonly user: ReadonlyArray<{
      readonly email: string;
      readonly id: string;
    }> | null | undefined;
    readonly winner: {
      readonly id: string;
    } | null | undefined;
  }> | null | undefined;
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
v1 = [
  (v0/*: any*/)
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
          "concreteType": "User",
          "kind": "LinkedField",
          "name": "winner",
          "plural": false,
          "selections": (v1/*: any*/),
          "storageKey": null
        },
        {
          "alias": null,
          "args": null,
          "concreteType": "User",
          "kind": "LinkedField",
          "name": "currentTurn",
          "plural": false,
          "selections": (v1/*: any*/),
          "storageKey": null
        },
        {
          "alias": null,
          "args": null,
          "concreteType": "User",
          "kind": "LinkedField",
          "name": "user",
          "plural": true,
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
  "type": "User",
  "abstractKey": null
};
})();

(node as any).hash = "2469cee25b2d949f01e7e45124e1f007";

export default node;
