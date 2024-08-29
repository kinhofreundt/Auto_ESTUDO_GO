import psycopg2

def test_connection():
    try:
        # Conectar ao banco de dados
        connection = psycopg2.connect(
            host="localhost",
            database="postgres",
            user="postgres",
            password="ericaserik"
        )

        # Criar um cursor para executar operações no banco de dados
        cursor = connection.cursor()

        # Executar um comando SQL para verificar a conexão
        cursor.execute("SELECT version();")
        version = cursor.fetchone()
        print("Conexão estabelecida com sucesso com o PostgreSQL:")
        print(version)

    except Exception as e:
        print("Erro ao conectar ao banco de dados:", e)

    finally:
        # Fechar a conexão e o cursor
        if connection:
            cursor.close()
            connection.close()
            print("Conexão com o banco de dados fechada.")

if __name__ == "__main__":
    test_connection()
