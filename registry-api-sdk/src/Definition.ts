export declare type Definition = {
    apis: APIMap,
    pages: Page[],
    components: ComponentMap,
    sidebars: SidebarMap,
}

export type Page = {
    name: number,
    path: string,
    title: string,
    components: PageComponent[],
}

export declare type API = {
    id: string,
    path: string,
    actions: APIActionMap,
}

export declare type APIMap = {
    [key: string]: API,
}

export declare type APIAction = {
    name: string,
    path: string,
    method: string,
}

export declare type APIActionMap = {
    [key: string]: APIAction,
}

export declare type Component = {
    name: string,
}

export declare type ComponentMap = {
    [key: string]: Component,
}

export declare type Sidebar = {
    name: string;
    links: SidebarLink[],
}

export declare type SidebarMap = {
    [key: string]: Sidebar,
}

export declare type SidebarLink = {
    pageName: string,
    title: string,
    children: SidebarLink[],
}

export declare type PageComponent = {
    id: string,
    baseComponentName: string
    config: Object,
}
