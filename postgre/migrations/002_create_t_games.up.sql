

CREATE TABLE IF NOT EXISTS gamelib.t_games
(
    id BIGSERIAL PRIMARY KEY,

    done BOOLEAN DEFAULT false,
    name TEXT NOT NULL UNIQUE,

    create_dt timestamptz NOT NULL DEFAULT now(),
    update_dt timestamptz NOT NULL DEFAULT now()
    );

CREATE FUNCTION gamelib.game_update_dt() RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
BEGIN
    new.update_dt = now();
RETURN NEW;
END;
$$;

CREATE TRIGGER game_update
    BEFORE UPDATE
    ON gamelib.t_games
    FOR EACH ROW
    EXECUTE PROCEDURE gamelib.game_update_dt();