INSERT INTO public.call (
	   agency
	 , latitude
	 , longitude
	 , incident 
	 , location
	 , datestamp
	 )  VALUES (
        _agency, 
	_latitude,
	_longitude,
	_incident,
	_location,
	_datestamp
        )
ON CONFLICT (agency,incident,datestamp) DO NOTHING;
-- agency, indcident and datestamp must be added as a constraint for this to work