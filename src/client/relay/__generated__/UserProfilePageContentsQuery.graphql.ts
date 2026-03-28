/**
 * @generated SignedSource<<17560edea1275b57a43d789bda80abd5>>
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
  readonly viewer: {
    readonly id: string;
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
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v2 = [
  (v1/*: any*/)
],
v3 = {
  "alias": null,
  "args": null,
  "concreteType": "User",
  "kind": "LinkedField",
  "name": "viewer",
  "plural": false,
  "selections": (v2/*: any*/),
  "storageKey": null
},
v4 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "id"
  }
],
v5 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "email",
  "storageKey": null
},
v6 = [
  (v1/*: any*/),
  (v5/*: any*/)
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "UserProfilePageContentsQuery",
    "selections": [
      (v3/*: any*/),
      {
        "alias": null,
        "args": (v4/*: any*/),
        "concreteType": "User",
        "kind": "LinkedField",
        "name": "user",
        "plural": false,
        "selections": [
          (v1/*: any*/),
          (v5/*: any*/),
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
      (v3/*: any*/),
      {
        "alias": null,
        "args": (v4/*: any*/),
        "concreteType": "User",
        "kind": "LinkedField",
        "name": "user",
        "plural": false,
        "selections": [
          (v1/*: any*/),
          (v5/*: any*/),
          {
            "alias": null,
            "args": null,
            "concreteType": "User",
            "kind": "LinkedField",
            "name": "friends",
            "plural": true,
            "selections": (v6/*: any*/),
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
                "selections": (v2/*: any*/),
                "storageKey": null
              },
              (v1/*: any*/),
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
                "selections": (v2/*: any*/),
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "concreteType": "User",
                "kind": "LinkedField",
                "name": "user",
                "plural": true,
                "selections": (v6/*: any*/),
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
    "cacheID": "dec9fc5ce7de55e3188b8c703ec3e0e1",
    "id": null,
    "metadata": {},
    "name": "UserProfilePageContentsQuery",
    "operationKind": "query",
    "text": "query UserProfilePageContentsQuery(\n  $id: ID!\n) {\n  viewer {\n    id\n  }\n  user(id: $id) {\n    id\n    email\n    ...UserProfileFriendsSection_user\n    ...UserProfileStats_user\n    ...UserProfileGamesSection_user\n  }\n}\n\nfragment UserProfileFriendsSection_user on User {\n  friends {\n    id\n    email\n  }\n}\n\nfragment UserProfileGamesSection_user on User {\n  id\n  games {\n    id\n    type\n    status\n    winner {\n      id\n    }\n    currentTurn {\n      id\n    }\n    user {\n      id\n      email\n    }\n  }\n}\n\nfragment UserProfileStats_user on User {\n  id\n  friends {\n    id\n  }\n  games {\n    winner {\n      id\n    }\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "6375f7c3deb5baca48e615b7be821be2";

export default node;
