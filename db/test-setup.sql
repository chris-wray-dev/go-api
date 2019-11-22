DROP DATABASE project_test;
CREATE DATABASE project_test;
\c project_test;

CREATE TABLE locations (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(30) UNIQUE NOT NULL, 
  text TEXT NOT NULL,
  image_url TEXT NOT NULL,
  lat FLOAT NOT NULL,
  long FLOAT NOT NULL,
  deleted_at TIMESTAMP
);

INSERT INTO locations (id, name, text, image_url, lat, long) 
VALUES (1, 'Hacienda', 'The legendary Manchester club, was originally conceived by Rob Gretton (Joy Division’s manager) and was largely funded by Tony Wilson’s Factory Records.
The Whitworth Street West venue, previously a Yacht showroom, opened in 1982 with a striking post-industrial interior designed by architect Ben Kelly.
Madonna made her first UK TV appearance live at the Hacienda in 1984.
However, the club struggled throughout it’s formative years - until 1986 - when it became one of the first UK clubs to play “house” music.
With the explosion of ecstasy-fuelled dance music that followed, The Hacienda became the home of house music in the UK in the late 1980s and a focal point of the “Madchester” scene… on the very same street that had seen the birth of the Northern Soul scene two decades earlier.
In the 1990s, drugs, gangs, shootings and badly managed finances all played a part in the club’s downfall
After a number of temporary closures, the Hacienda finally closed it’s doors permanently in 1997, and the building was demolished to make way for an apartment block – “The Hacienda Apartments” - in 2002.',
'url tbc',
2.321,
4.45654
);

INSERT INTO locations (id, name, text, image_url, lat, long) 
VALUES (2, 'The Twisted Wheel', 'Famed for the birth of the Northern Soul scene.
The original Twisted Wheel club opened in 1963 as a blues & soul live music venue at it’s first location on Brazennose Street.
The legendary Whitworth Street site opened on 18th September 1965, with The Spencer Davis Group headlining the opening night.
Twisted Wheel DJs famously imported huge amounts of rare soul records from the US – records that were rare even in the US.
All-night club sessions took place every Saturday night, from 11pm until 7.30am.
This was the birth of the ‘Northern Soul’ scene – and similar venues later appeared  across the North West - Wigan Casino and the Blackpool Mecca.
Live artists – which included Edwin Starr, Mary Wells and Tina Turner – tended to come on stage at 2am!
Sadly the original building was demolished in 2013 to make way for a hotel.
However “The Wheel” has now re-located to Night People (105-107 Princess St, Manchester, M1 6DD) where regular soul and R&B sessions, along with memorabilia from the original club, re-create the storming times of the Whitworth Street club.
https://www.nightpeoplemcr.com/
nightpeoplemcr.com
Night People
Official Website for Night People . Find all Upcoming Events with Lineups and set times, latest news and announcements and contact information.',
'url tbc',
4.321,
6.587756
);

INSERT INTO locations (id, name, text, image_url, lat, long) 
VALUES (
  3, 
  'The Free Trade Hall', 
  'This Grade II listed building was constructed as a public hall (1853 – 1856) on the site of the Peterloo Massacre and became a concert hall and home to the Halle Orchestra in 1858.
It was re-built following a direct hit during a blitz bombing campaign in 1940.
In the 1960s and 1970s it was home to two of the most talked about moments in Western popular music history.
On 17th May 1966, one member of a particularly disgruntled audience heckled Bob Dylan - ‘Judas!’ –  which was followed by ensuing cheers from the crowd, after the folk artist abandoned his trademark folk acoustic guitar and went electric with full band… “I don’t believe you… you’re a liar!” Dylan retorted to the crowd, before turning to his band…  “Play loud!“.
Nearly exactly 10 years later, on 4th June 1976, local Manchester band The Buzzcocks put on a gig at the hall supporting an unknown band from London – known as ‘The Sex Pistols’.
This first Sex Pistols Manchester gig – pre-fame and pre-notoriety - was famously poorly attended, but for those who did attend - future members of Joy Division / New Order, The Fall, The Smiths, Magazine and Factory Records kingpins, label boss Tony Wilson and producer Martin Hannett – this signalled the start of the Punk era and was a catalyst for pretty much everything that happened in Manchester Music in the ensuing decades.
The Sex Pistols played the Free Trade Hall again just over a month later - on 20th July 1976 – this time to a packed out venue.',
'url tbc',
5.321,
7.4563456
);

-- SELECT * FROM locations