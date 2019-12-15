CREATE TABLE notification (
  id           BIGSERIAL,
  created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  subject      TEXT                     NOT NULL,
  notification_type TEXT                     NOT NULL,
  from_address         TEXT                     NULL,
  body_html    TEXT             NULL,
  body_plain_text         TEXT   NULL,
  CONSTRAINT pk_notification PRIMARY KEY (id),
  CHECK (body_html IS NOT NULL OR body_plain_text IS NOT NULL)
);

CREATE TABLE notification_recipient (
  id           BIGSERIAL ,
  notification_id   BIGINT NOT NULL,
  recipient_type TEXT   NOT NULL,
  name         TEXT,
  address      TEXT NOT NULL,
  CONSTRAINT pk_notification_recipient PRIMARY KEY (id),
  CONSTRAINT fk_notification_recipient__notification FOREIGN KEY (notification_id) REFERENCES notification (id)
);


CREATE TABLE notification_attachment (
  id         BIGSERIAL,
  notification_id BIGINT NOT NULL,
  name       TEXT NOT NULL,
  data       BYTEA NOT NULL,
  CONSTRAINT pk_notification_attachment PRIMARY KEY (id),
  CONSTRAINT fk_notification_attachment__notification FOREIGN KEY (notification_id) REFERENCES notification (id)
);
