BEGIN;

CREATE TABLE alembic_version (
    version_num VARCHAR(32) NOT NULL, 
    CONSTRAINT alembic_version_pkc PRIMARY KEY (version_num)
);

-- Running upgrade  -> 51638789f439

CREATE TABLE town_node (
    id SERIAL NOT NULL, 
    pub_key TEXT NOT NULL, 
    amount_collected INTEGER NOT NULL, 
    amount_received INTEGER NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP WITH TIME ZONE, 
    PRIMARY KEY (id)
);

ALTER TABLE town_node ADD CONSTRAINT uq_node_pub_key UNIQUE (pub_key);

CREATE TABLE town_address (
    id SERIAL NOT NULL, 
    value VARCHAR(35) NOT NULL, 
    amount_collected INTEGER NOT NULL, 
    amount_received INTEGER NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP WITH TIME ZONE, 
    PRIMARY KEY (id)
);

ALTER TABLE town_address ADD CONSTRAINT uq_address_value UNIQUE (value);

CREATE TABLE town_slug (
    id SERIAL NOT NULL, 
    slug VARCHAR(100) NOT NULL, 
    current_id INTEGER, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP WITH TIME ZONE, 
    PRIMARY KEY (id), 
    CONSTRAINT town_slug_current_id_fkey FOREIGN KEY(current_id) REFERENCES town_slug (id)
);

CREATE INDEX town_article_current_id ON town_slug (current_id);

ALTER TABLE town_slug ADD CONSTRAINT uq_slug UNIQUE (slug);

CREATE TABLE town_article (
    id SERIAL NOT NULL, 
    title VARCHAR(100) NOT NULL, 
    slug VARCHAR(100) NOT NULL, 
    lang INTEGER NOT NULL, 
    amount_collected INTEGER NOT NULL, 
    amount_received INTEGER NOT NULL, 
    status INTEGER NOT NULL, 
    subtitle VARCHAR(255), 
    body TEXT NOT NULL, 
    address_id INTEGER, 
    node_id INTEGER, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP WITH TIME ZONE, 
    published_at TIMESTAMP WITH TIME ZONE, 
    PRIMARY KEY (id), 
    CONSTRAINT town_address_id_fkey FOREIGN KEY(address_id) REFERENCES town_address (id), 
    CONSTRAINT town_node_id_fkey FOREIGN KEY(node_id) REFERENCES town_node (id)
);

CREATE INDEX town_article_address_id ON town_article (address_id);

CREATE INDEX town_article_node_id ON town_article (node_id);

CREATE INDEX town_article_slug ON town_article (slug);

ALTER TABLE town_article ADD CONSTRAINT uq_article_slug UNIQUE (slug);

CREATE TABLE town_order (
    id SERIAL NOT NULL, 
    public_id VARCHAR(255) NOT NULL, 
    description TEXT NOT NULL, 
    amount INTEGER NOT NULL, 
    status INTEGER NOT NULL, 
    fee INTEGER NOT NULL, 
    fiat_value FLOAT NOT NULL, 
    currency INTEGER NOT NULL, 
    notes TEXT NOT NULL, 
    payreq TEXT NOT NULL, 
    charge_created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    charge_settle_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    updated_at TIMESTAMP WITH TIME ZONE, 
    claimed_at TIMESTAMP WITH TIME ZONE, 
    PRIMARY KEY (id)
);

CREATE INDEX town_order_public_id ON town_order (public_id);

CREATE TABLE town_article_reaction (
    id SERIAL NOT NULL, 
    article_id INTEGER NOT NULL, 
    order_id INTEGER NOT NULL, 
    emoji VARCHAR(30) NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    PRIMARY KEY (id)
);

CREATE INDEX town_article_reaction_article_id ON town_article_reaction (article_id);

CREATE TABLE town_article_comment (
    id SERIAL NOT NULL, 
    article_id INTEGER NOT NULL, 
    order_id INTEGER NOT NULL, 
    body TEXT NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, 
    PRIMARY KEY (id)
);

CREATE INDEX town_article_comment_article_id ON town_article_comment (article_id);

INSERT INTO alembic_version (version_num) VALUES ('51638789f439');

COMMIT;

