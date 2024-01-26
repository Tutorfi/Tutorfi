package storage
import(
	_ "database/sql"
)
type PostgresError struct{
	message string
	sqlError error//idk the format of sql errors, so leaving this here in case we need to access the raw error

}

func (se *PostgresError) Error() (string){
	return se.message
}

func(se *PostgresError) SQLError(){
	return sqlError
}