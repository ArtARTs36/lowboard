// @ts-ignore
import fetch from "isomorphic-unfetch";

import {type Definition} from "./Definition";

type Config = {
    baseUrl: string;
};

// @ts-ignore
export class Client {
    private readonly baseUrl: string;

    constructor(config: Config) {
        this.baseUrl = config.baseUrl;
    }

    public getDefinition(): Promise<Definition> {
        return this.request<Definition>(`/definition`);
    }

    private async request<T>(endpoint: string, options?: RequestInit): Promise<T> {
        const url = `${this.baseUrl}${endpoint}`;
        const headers = {
            'Content-Type': 'application/json',
        };
        const config = {
            ...options,
            headers,
        };

        return fetch(url, config)
            // When got a response call a `json` method on it
            .then((response) => response.json())
            // and return the result data.
            .then((data) => data as T);
    }
}
