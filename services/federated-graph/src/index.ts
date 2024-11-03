import { ApolloServer } from "@apollo/server";
import { ApolloGateway, IntrospectAndCompose } from "@apollo/gateway";
import { startStandaloneServer } from "@apollo/server/standalone";

const INV_SERVICE_PORT = process.env.INV_SERVICE_PORT;
const PRODUCT_SERVICE_PORT = process.env.PRODUCT_SERVICE_PORT;

const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      { name: "inventory", url: `http://localhost:${INV_SERVICE_PORT}/query` },
      {
        name: "product",
        url: `http://localhost:${PRODUCT_SERVICE_PORT}/query`,
      },
    ],
  }),
});

const server = new ApolloServer({ gateway });

const { url } = await startStandaloneServer(server, {
  listen: { port: 4000 },
});
console.log(`🚀 Federated Gateway ready at ${url}`);
