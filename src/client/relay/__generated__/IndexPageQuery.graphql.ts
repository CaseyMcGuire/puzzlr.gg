/**
 * @generated SignedSource<<da07a289a476701897679e8b17008cce>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type IndexPageQuery$variables = Record<PropertyKey, never>;
export type IndexPageQuery$data = {
  readonly todos: ReadonlyArray<{
    readonly id: string;
  }>;
};
export type IndexPageQuery = {
  response: IndexPageQuery$data;
  variables: IndexPageQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "Todo",
    "kind": "LinkedField",
    "name": "todos",
    "plural": true,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "id",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "IndexPageQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "IndexPageQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "cabaa9183b127418eb84b5ce1ddfc1cf",
    "id": null,
    "metadata": {},
    "name": "IndexPageQuery",
    "operationKind": "query",
    "text": "query IndexPageQuery {\n  todos {\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "d9ab5049e55354d90509d489a78fe464";

export default node;
