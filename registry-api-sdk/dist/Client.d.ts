import { type Definition } from "./Definition";
type Config = {
    baseUrl: string;
};
export declare class Client {
    private readonly baseUrl;
    constructor(config: Config);
    getDefinition(): Promise<Definition>;
    private request;
}
export {};
//# sourceMappingURL=Client.d.ts.map