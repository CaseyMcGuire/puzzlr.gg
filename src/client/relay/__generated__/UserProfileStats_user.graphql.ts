/**
 * @generated SignedSource<<47f38aca6e586d58dbc41dd37a623934>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type UserProfileStats_user$data = {
  readonly friends: ReadonlyArray<{
    readonly id: string;
  }> | null | undefined;
  readonly games: ReadonlyArray<{
    readonly winnerPlayer: {
      readonly user: {
        readonly id: string;
      } | null | undefined;
    } | null | undefined;
  }>;
  readonly id: string;
  readonly " $fragmentType": "UserProfileStats_user";
};
export type UserProfileStats_user$key = {
  readonly " $data"?: UserProfileStats_user$data;
  readonly " $fragmentSpreads": FragmentRefs<"UserProfileStats_user">;
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
  "name": "UserProfileStats_user",
  "selections": [
    (v0/*: any*/),
    {
      "alias": null,
      "args": null,
      "concreteType": "User",
      "kind": "LinkedField",
      "name": "friends",
      "plural": true,
      "selections": (v1/*: any*/),
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "concreteType": "Game",
      "kind": "LinkedField",
      "name": "games",
      "plural": true,
      "selections": [
        {
          "alias": null,
          "args": null,
          "concreteType": "GamePlayer",
          "kind": "LinkedField",
          "name": "winnerPlayer",
          "plural": false,
          "selections": [
            {
              "alias": null,
              "args": null,
              "concreteType": "User",
              "kind": "LinkedField",
              "name": "user",
              "plural": false,
              "selections": (v1/*: any*/),
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

(node as any).hash = "bd591652d43b136e64607faf791f9cbd";

export default node;
