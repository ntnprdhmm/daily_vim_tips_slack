CREATE DATABASE IF NOT EXISTS daily_vim
CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

USE daily_vim;

CREATE TABLE IF NOT EXISTS tips (
  `command` VARCHAR(100) NOT NULL,
  `description` TEXT NOT NULL,
  `index` INT NOT NULL UNIQUE AUTO_INCREMENT,
  `posted` TINYINT DEFAULT 0,
  PRIMARY KEY (`command`)
) ENGINE = InnoDB;

INSERT INTO tips (command, description)
VALUES
(
  "[count]dd",
  "Delete the next `count` lines from the line where the cursor is.\n\n- `dd` or `1dd` will remove the current line\n- `4dd` will remove the current line and the next three ones"
),
(
  "[count]yy",
  "Copy the next `count` lines from the line where the cursor is.\n\n- `yy` or `1yy` will copy the current line\n- `3yy` will copy the current line and the 2 next ones\n\nThen, to paste the copied lines, use `p`\n\n- `p` or `1p` will insert the copied lines under the current line\n- `5p` will insert the copied lines 5 times under the current line"
),
(
  "[count]X",
  "Delete the `count` charaters before the cursor, on the same line.\n\n- `X` or `1X` will remove the previous character.\n\n- `4X` will remove the 4 previous characters. If there are less than 4 characters, it will *not* erase characters of the previous line."
),
(
  "[count]r[char]",
  "Replace the `count` characters at the right of the cursor by `char`.\n\n- `rb` or `1rb` will replace the char at the right of the cursor by `b`\n- `3rb` will replace the 3 characters at the right of the cursor by `b`\n"
),
(
  "[count]f[char]",
  "Put the cursor just before the `count`'th occurence of `char` after the current position of the cursor, in the current line.\n\n*line example*: `elle a le #regard qui tue, tchikita`\n(With the cursor `#` just before `regard`)\n- `ft` = `elle a le regard qui #tue, tchikita`\n- `fz` = `elle a le #regard qui tue, tchikita`\n- `3fi` = `elle a le regard qui tue, tchik#ita`"
),
(
  "<",
  "shift the text of the line to the left (~= remove a tab)"
),
(
  "[count]G",
  "move the cursor at the start of the line `count`.\n(line numbers start at 1, not 0)"
),
(
  "gg",
  "move the cursor at the start of the first line of the document (the line 1)"
),
(
  "[count]t[char]",
  "move the cursor before the character just before `char` in the current line.\n\n*line example*: `elle a le #regard qui tue, tchikita`\n(With the cursor `#` just before `regard`)\n- `tu` = `elle a le regard qui #tue, tchikita`\n- `tz` = `elle a le #regard qui tue, tchikita\n- `2te` = `elle a le #regard qui t#ue, tchikita``"
),
(
  "G",
  "move the cursor at the start of the last line of the document"
),
(
  ">",
  "shift the text of the line to the right (add a tab)"
),
(
  "[count]T[char]",
  "move the cursor just after the previous occurence of `char` in the current line.\n\n*line example*: `elle a le #regard qui tue, tchikita`\n(With the cursor `#` just before `regard`)\n- `Tl` = `elle a l#e regard qui tue, tchikita`\n- `Tz` = `elle a le #regard qui tue, tchikita\n- `2Tz` = `elle# a le regard qui tue, tchikita`"
),
(
  "[count]F[char]",
  "Put the cursor just before the `count`'th occurence of `char` before the current position of the cursor, in the current line.\n\n*line example*: `elle a le #regard qui tue, tchikita`\n(With the cursor `#` just before `regard`)\n- `2Fe` = `ell#e a le regard qui tue, tchikita`\n- `fz` = `elle a le #regard qui tue, tchikita`"
);