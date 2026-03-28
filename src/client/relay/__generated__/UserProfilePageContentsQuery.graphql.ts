/**
 * @generated SignedSource<<ff50c6e42ce63e4bff821ac40df4521d>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type ViewerFriendshipStatus = "FRIENDS" | "NOT_APPLICABLE" | "NOT_FRIENDS" | "NOT_LOGGED_IN" | "REQUEST_RECEIVED" | "REQUEST_SENT" | "%future added value";
export type UserProfilePageContentsQuery$variables = {
  id: string;
};
export type UserProfilePageContentsQuery$data = {
  readonly user: {
    readonly email: string;
    readonly id: string;
    readonly viewerFriendshipStatus: ViewerFriendshipStatus;
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
v4 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "viewerFriendshipStatus",
  "storageKey": null
},
v5 = [
  (v2/*: any*/),
  (v3/*: any*/)
],
v6 = [
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
          (v4/*: any*/),
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
          (v4/*: any*/),
          {
            "alias": null,
            "args": null,
            "concreteType": "User",
            "kind": "LinkedField",
            "name": "friends",
            "plural": true,
            "selections": (v5/*: any*/),
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
                "selections": (v6/*: any*/),
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
                "selections": (v6/*: any*/),
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "concreteType": "User",
                "kind": "LinkedField",
                "name": "user",
                "plural": true,
                "selections": (v5/*: any*/),
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
    "cacheID": "171b6a59a9035a66a5e226670fe25875",
    "id": null,
    "metadata": {},
    "name": "UserProfilePageContentsQuery",
    "operationKind": "query",
    "text": "query UserProfilePageContentsQuery(\n  $id: ID!\n) {\n  user(id: $id) {\n    id\n    email\n    viewerFriendshipStatus\n    ...UserProfileFriendsSection_user\n    ...UserProfileStats_user\n    ...UserProfileGamesSection_user\n  }\n}\n\nfragment UserProfileFriendsSection_user on User {\n  friends {\n    id\n    email\n  }\n}\n\nfragment UserProfileGamesSection_user on User {\n  id\n  games {\n    id\n    type\n    status\n    winner {\n      id\n    }\n    currentTurn {\n      id\n    }\n    user {\n      id\n      email\n    }\n  }\n}\n\nfragment UserProfileStats_user on User {\n  id\n  friends {\n    id\n  }\n  games {\n    winner {\n      id\n    }\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "85215a5dc4dbc3886aa3cbd152b5015d";

export default node;
