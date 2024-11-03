import { ApolloServer } from "@apollo/server";
import { ApolloGateway, IntrospectAndCompose } from "@apollo/gateway";
import { startStandaloneServer } from "@apollo/server/standalone";

const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [{ name: "inventory", url: "http://localhost:8081/query" }],
  }),
});

const server = new ApolloServer({ gateway });

const { url } = await startStandaloneServer(server, {
  listen: { port: 4000 },
});
console.log(`🚀 Federated Gateway ready at ${url}`);
