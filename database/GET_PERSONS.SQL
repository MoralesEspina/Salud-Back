SELECT 
    p.uuid, 
    p.fullname, 
    p.cui,
    p.sueldo,
    p.admissionDate,
    j.name as job,
    j.uuid as uuid_job,
    w.name as workdependency,
    w.uuid as uuid_workdependency,
    esp.name as especiality,
    esp.uuid as uuid_especiality,
    reu.name as reubication,
    reu.uuid as uuid_reubication
    FROM person p
    INNER JOIN job w ON p.Workdependency = w.uuid
	LEFT JOIN job j ON p.job = j.uuid
    LEFT JOIN job esp ON p.especiality = esp.uuid
    LEFT JOIN job reu ON p.reubication = reu.uuid