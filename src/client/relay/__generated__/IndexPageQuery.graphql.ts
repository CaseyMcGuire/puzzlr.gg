/**
 * @generated SignedSource<<376dd785695456fb410525a68b54dcc6>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type IndexPageQuery$variables = Record<PropertyKey, never>;
export type IndexPageQuery$data = {
  readonly users: ReadonlyArray<{
    readonly games: ReadonlyArray<{
      readonly id: string;
    }> | null | undefined;
  }>;
};
export type IndexPageQuery = {
  response: IndexPageQuery$data;
  variables: IndexPageQuery$variables;
};

const node: ConcreteRequest = (function(){
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
  "concreteType": "Game",
  "kind": "LinkedField",
  "name": "games",
  "plural": true,
  "selections": [
    (v0/*: any*/)
  ],
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "IndexPageQuery",
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "User",
        "kind": "LinkedField",
        "name": "users",
        "plural": true,
        "selections": [
          (v1/*: any*/)
        ],
        "storageKey": null
      }
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "IndexPageQuery",
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "User",
        "kind": "LinkedField",
        "name": "users",
        "plural": true,
        "selections": [
          (v1/*: any*/),
          (v0/*: any*/)
        ],
        "storageKey": null
      }
    ]
  },
  "params": {
    "cacheID": "421df8807cce3cf54a05bb3bb95c8b03",
    "id": null,
    "metadata": {},
    "name": "IndexPageQuery",
    "operationKind": "query",
    "text": "query IndexPageQuery {\n  users {\n    games {\n      id\n    }\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "32d1609b23e1928758a95a11261eecc3";

export default node;
