/**
 * @generated SignedSource<<7e27625732634ddcf712d3e37b9a80cb>>
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
    readonly winner: {
      readonly id: string;
    } | null | undefined;
  }> | null | undefined;
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
          "concreteType": "User",
          "kind": "LinkedField",
          "name": "winner",
          "plural": false,
          "selections": (v1/*: any*/),
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

(node as any).hash = "06e1fa4b4aa9c24eef9b1561dbb2ff83";

export default node;
