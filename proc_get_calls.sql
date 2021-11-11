-- FUNCTION: public.get_recent_calls()

-- DROP FUNCTION public.get_recent_calls();

CREATE OR REPLACE FUNCTION public.get_recent_calls(
	)
    RETURNS TABLE(id bigint, agency text, latitude text, longitude text, incident text, location text, reported_date date) 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
begin
	return query 
		select
        call.id, 
        call.agency, 
        call.latitude, 
        call.longitude, 
        call.incident, 
        call.location, 
        call.datestamp
	FROM public.call
    order by
            datestamp DESC;
	end;
$BODY$;

ALTER FUNCTION public.get_recent_calls()
    OWNER TO pi;
