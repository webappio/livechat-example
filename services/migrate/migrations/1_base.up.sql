CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE channels(
    uuid uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

    name text UNIQUE NOT NULL,
    topic text NOT NULL,
    description text NOT NULL
);

CREATE OR REPLACE FUNCTION channels_notify()
    RETURNS trigger AS
$$
BEGIN
    PERFORM pg_notify(
                    'channels',
                    COALESCE(NEW.uuid, OLD.uuid)::text
                );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER channels
    AFTER INSERT OR DELETE OR UPDATE
    ON channels
    FOR EACH ROW
EXECUTE PROCEDURE channels_notify();


CREATE TABLE users(
    uuid uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

    name text UNIQUE NOT NULL,
    password_hash text NOT NULL
);


CREATE OR REPLACE FUNCTION users_notify()
    RETURNS trigger AS
$$
BEGIN
    PERFORM pg_notify(
                    'users',
                    COALESCE(NEW.uuid, OLD.uuid)::text
                );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER users
    AFTER INSERT OR DELETE OR UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE users_notify();


CREATE TABLE channel_messages(
    channel_uuid uuid NOT NULL REFERENCES channels(uuid),
    user_uuid uuid NOT NULL REFERENCES users(uuid),

    index int NOT NULL,
    text text NOT NULL,
    time timestamp NOT NULL DEFAULT NOW(),

    PRIMARY KEY (channel_uuid, index)
);

CREATE OR REPLACE FUNCTION channel_messages_notify()
    RETURNS trigger AS
$$
BEGIN
    PERFORM pg_notify(
                    'channel_messages',
                    COALESCE(NEW.channel_uuid, OLD.channel_uuid)::text || ',' || COALESCE(NEW.index, OLD.index)
                );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER channel_messages
    AFTER INSERT OR DELETE OR UPDATE
    ON channel_messages
    FOR EACH ROW
EXECUTE PROCEDURE channel_messages_notify();

CREATE TABLE direct_messages(
    from_user_uuid uuid NOT NULL REFERENCES users(uuid),
    to_user_uuid uuid NOT NULL REFERENCES users(uuid),

    index int NOT NULL,
    text text NOT NULL,

    PRIMARY KEY(to_user_uuid, index)
);

CREATE OR REPLACE FUNCTION direct_messages_notify()
    RETURNS trigger AS
$$
BEGIN
    PERFORM pg_notify(
                    'direct_messages',
                    COALESCE(NEW.to_user_uuid, OLD.to_user_uuid)::text || ',' || COALESCE(NEW.index, OLD.index)
                );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER direct_messages
    AFTER INSERT OR DELETE OR UPDATE
    ON direct_messages
    FOR EACH ROW
EXECUTE PROCEDURE direct_messages_notify();