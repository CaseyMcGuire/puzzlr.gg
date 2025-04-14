import {Environment, Network, RecordSource, RequestParameters, Store, Variables} from "relay-runtime";
//import CsrfUtils from "../utils/CsrfUtils";

export class RelayConfig {
  static getEnvironment(): Environment {
    return new Environment({
      network: Network.create(this.fetchQuery),
      store: new Store(new RecordSource())
    });

  }

  private static async fetchQuery(
    operation: RequestParameters,
    variables: Variables
  ) {
    //const csrfToken = CsrfUtils.getToken();
    //const csrfHeader = CsrfUtils.getHeader();
    const response = await fetch("/graphql", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        query: operation.text,
        variables
      })
    });
    return await response.json();
  }
}