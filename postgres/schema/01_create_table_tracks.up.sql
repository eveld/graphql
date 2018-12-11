CREATE TABLE tracks (
  id            VARCHAR(255) PRIMARY KEY,
  slug          VARCHAR(255) NOT NULL,
  title         VARCHAR(255) NOT NULL,
  teaser        VARCHAR(255) NOT NULL,
  description   TEXT NOT NULL,
  deleted       BIGINT NOT NULL DEFAULT 0
);

CREATE TRIGGER trigger_tracks_genid BEFORE INSERT ON tracks FOR EACH ROW EXECUTE PROCEDURE unique_short_id();