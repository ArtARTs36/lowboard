// @ts-ignore
import {Client} from "./Client";

// @ts-ignore
export class SDK {
    client: Client;

    constructor(config: { baseUrl: string }) {
        this.client = new Client(config);
    }
}
