import {GRAPHQL_ENDPOINT} from '../const';

export function gql(query, variables = {}) {
  return fetch(GRAPHQL_ENDPOINT, {
    method: 'POST',
    body: JSON.stringify({query, variables}),
  })
    .then(resp => resp.json())
    .then(resp => resp.data);
}
