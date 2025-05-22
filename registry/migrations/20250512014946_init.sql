-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages (
    name VARCHAR NOT NULL PRIMARY KEY,
    path VARCHAR NOT NULL UNIQUE,
    title VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE page_components (
    id VARCHAR NOT NULL,

    page_name VARCHAR NOT NULL,
    base_component_name VARCHAR NOT NULL,
    config JSONB NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT page_components_pk PRIMARY KEY (id),
    CONSTRAINT page_components_page_name_fk FOREIGN KEY (page_name) REFERENCES pages(name),
    CONSTRAINT page_components_page_base_component_name_fk FOREIGN KEY (base_component_name) REFERENCES components(name)
);

CREATE TABLE components (
    name VARCHAR NOT NULL,
    title VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT components_pk PRIMARY KEY (name)
);

INSERT INTO components (name, title, created_at)
VALUES
('table', 'Table', '2025-05-17 01:41:00'),
('table-page', 'Table Page', '2025-05-17 01:41:00');

CREATE TABLE apis (
    id VARCHAR NOT NULL,
    path VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT apis_pk PRIMARY KEY (id)
);

CREATE TABLE api_actions (
    name VARCHAR NOT NULL,
    api_id VARCHAR NOT NULL,
    method VARCHAR NOT NULL,
    path VARCHAR NOT NULL,
    description VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT api_actions_pk PRIMARY KEY (name, api_id),
    CONSTRAINT api_actions_api_id_fk FOREIGN KEY (api_id) REFERENCES apis(id)
);

CREATE TABLE sidebars (
    name VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT sidebars_pk PRIMARY KEY (name)
);

CREATE TABLE sidebar_links (
    id VARCHAR NOT NULL, -- uuid

    sidebar_name VARCHAR NOT NULL,
    page_name VARCHAR NOT NULL,
    title VARCHAR NOT NULL,

    parent_id VARCHAR,
    icon VARCHAR,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT sidebar_links_pk PRIMARY KEY (id),
    CONSTRAINT sidebar_links_page_fk FOREIGN KEY (page_name) REFERENCES pages (name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE page_components;
DROP TABLE components;
DROP TABLE pages;
DROP TABLE api_actions;
DROP TABLE apis;
DROP TABLE sidebar_links;
DROP TABLE sidebars;
-- +goose StatementEnd
