# 記事データ2つ
INSERT INTO articles (
  title,
  contents,
  username,
  nice,
  created_at
) VALUES ( 
  "firstPost", "This is my first blog", "saki", 2, now()
);

INSERT INTO articles (
  title,
  contents,
  username,
  nice,
  created_at
) VALUES ( 
  "2nd", "This is my second blog", "saki", 4, now()
);

INSERT INTO comments (
  article_id,
  message,
  created_at
) VALUES (
1, "1st comment yeah", now()
);

INSERT INTO comments (
  article_id,
  message
) VALUES (
1, "welcome"
);
