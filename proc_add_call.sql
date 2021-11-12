-- PROCEDURE: public.add_call(text, text, text, text, text, date)

-- DROP PROCEDURE public.add_call(text, text, text, text, text, date);

CREATE OR REPLACE PROCEDURE public.add_call(
	_agency text,
	_latitude text,
	_longitude text,
	_incident text,
	_location text,
	_call_date date,
	_call_time time)
LANGUAGE 'sql'
AS $BODY$
INSERT INTO public.call (
	   agency
	 , latitude
	 , longitude
	 , incident 
	 , location
	 , call_date
	, call_time
	 )  VALUES (
        _agency, 
	_latitude,
	_longitude,
	_incident,
	_location,
	_call_date,
		 _call_time
        )
ON CONFLICT (agency,incident,call_date, call_time) DO NOTHING;
$BODY$;
