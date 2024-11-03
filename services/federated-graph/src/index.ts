import { ApolloServer } from "@apollo/server";
import { ApolloGateway, IntrospectAndCompose } from "@apollo/gateway";
import { startStandaloneServer } from "@apollo/server/standalone";

const INV_HOST_NAME = process.env.INV_HOST_NAME;
const PRODUCT_HOST_NAME = process.env.PRODUCT_HOST_NAME;

const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      {
        name: "inventory",
        url: `http://${INV_HOST_NAME}:8080/query`,
      },
      {
        name: "product",
        url: `http://${PRODUCT_HOST_NAME}:8080/query`,
      },
    ],
  }),
});

const server = new ApolloServer({ gateway });

const { url } = await startStandaloneServer(server, {
  listen: { port: 4000 },
});
console.log(`🚀 Federated Gateway ready at ${url}`);
