/**
 * @generated SignedSource<<29392a6321c3892cc28c1836c7bd6a9d>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type UserProfilePageContentsQuery$variables = {
  id: string;
};
export type UserProfilePageContentsQuery$data = {
  readonly user: {
    readonly email: string;
    readonly id: string;
    readonly " $fragmentSpreads": FragmentRefs<"UserProfileFriendsSection_user" | "UserProfileGamesSection_user" | "UserProfileStats_user">;
  } | null | undefined;
};
export type UserProfilePageContentsQuery = {
  response: UserProfilePageContentsQuery$data;
  variables: UserProfilePageContentsQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "id"
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "id"
  }
],
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v3 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "email",
  "storageKey": null
},
v4 = [
  (v2/*: any*/),
  (v3/*: any*/)
],
v5 = [
  (v2/*: any*/)
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "UserProfilePageContentsQuery",
    "selections": [
      {
        "alias": null,
        "args": (v1/*: any*/),
        "concreteType": "User",
        "kind": "LinkedField",
        "name": "user",
        "plural": false,
        "selections": [
          (v2/*: any*/),
          (v3/*: any*/),
          {
            "args": null,
            "kind": "FragmentSpread",
            "name": "UserProfileFriendsSection_user"
          },
          {
            "args": null,
            "kind": "FragmentSpread",
            "name": "UserProfileStats_user"
          },
          {
            "args": null,
            "kind": "FragmentSpread",
            "name": "UserProfileGamesSection_user"
          }
        ],
        "storageKey": null
      }
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "UserProfilePageContentsQuery",
    "selections": [
      {
        "alias": null,
        "args": (v1/*: any*/),
        "concreteType": "User",
        "kind": "LinkedField",
        "name": "user",
        "plural": false,
        "selections": [
          (v2/*: any*/),
          (v3/*: any*/),
          {
            "alias": null,
            "args": null,
            "concreteType": "User",
            "kind": "LinkedField",
            "name": "friends",
            "plural": true,
            "selections": (v4/*: any*/),
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
                "selections": (v5/*: any*/),
                "storageKey": null
              },
              (v2/*: any*/),
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
                "name": "currentTurn",
                "plural": false,
                "selections": (v5/*: any*/),
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "concreteType": "User",
                "kind": "LinkedField",
                "name": "user",
                "plural": true,
                "selections": (v4/*: any*/),
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ]
  },
  "params": {
    "cacheID": "ccb22a49527f79fa147d5e2a67da2c8b",
    "id": null,
    "metadata": {},
    "name": "UserProfilePageContentsQuery",
    "operationKind": "query",
    "text": "query UserProfilePageContentsQuery(\n  $id: ID!\n) {\n  user(id: $id) {\n    id\n    email\n    ...UserProfileFriendsSection_user\n    ...UserProfileStats_user\n    ...UserProfileGamesSection_user\n  }\n}\n\nfragment UserProfileFriendsSection_user on User {\n  friends {\n    id\n    email\n  }\n}\n\nfragment UserProfileGamesSection_user on User {\n  id\n  games {\n    id\n    type\n    status\n    winner {\n      id\n    }\n    currentTurn {\n      id\n    }\n    user {\n      id\n      email\n    }\n  }\n}\n\nfragment UserProfileStats_user on User {\n  id\n  friends {\n    id\n  }\n  games {\n    winner {\n      id\n    }\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "f2664f840bb53a7fae575bfce9c6f400";

export default node;
