CREATE TABLE challenges (
  id            VARCHAR(255) PRIMARY KEY,
  slug          VARCHAR(255) NOT NULL,
  title         VARCHAR(255) NOT NULL,
  teaser        VARCHAR(255) NOT NULL,
  assignment    TEXT NOT NULL,
  difficulty    VARCHAR(255) NOT NULL DEFAULT 'basic',
  timelimit     INT NOT NULL DEFAULT 300,
  deleted       BIGINT NOT NULL DEFAULT 0
);

CREATE TRIGGER trigger_challenges_genid BEFORE INSERT ON challenges FOR EACH ROW EXECUTE PROCEDURE unique_short_id();