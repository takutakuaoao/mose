package app_test

import (
	"database/sql"
	"github.com/go-playground/assert/v2"
	_ "github.com/lib/pq"
	"github.com/takutakuaoao/mose/api/app"
	"github.com/takutakuaoao/mose/api/app/mock"
	"github.com/takutakuaoao/mose/api/test"
	"os"
	"testing"
)

func Test_プロジェクト登録(t *testing.T) {
	t.Run("登録の成功時", func(t *testing.T) {
		id := int64(1)

		type ExpectedTitle = string
		type ExpectedLoginUserID = int64

		makeMock := func() (*mock.ProjectRepositoryMock, ExpectedTitle, ExpectedLoginUserID) {
			return &mock.ProjectRepositoryMock{
					RegisterFunc: func(title string, loginUserID int64) int64 {
						return id
					},
				},
				"Test",
				2
		}

		repository, expectedTitle, expectedLoginUserID := makeMock()

		t.Run("リポジトリ登録処理がタイトルとログインユーザーのIDと共に１度呼ばれる", func(t *testing.T) {
			register := app.NewProjectRegister(repository)

			register.Register(expectedTitle, expectedLoginUserID)

			assert_リポジトリの登録処理が1度だけタイトルとログインユーザーIDを受け取る(repository, expectedTitle, expectedLoginUserID, t)
		})

		t.Run("新規IDが返される", func(t *testing.T) {
			register := app.NewProjectRegister(repository)

			result := register.Register(expectedTitle, expectedLoginUserID)

			assert.Equal(t, id, result)
		})
	})
}

func assert_リポジトリの登録処理が1度だけタイトルとログインユーザーIDを受け取る(repository *mock.ProjectRepositoryMock, title string, loginUserID int64, t *testing.T) {
	assert.Equal(t, 1, len(repository.RegisterCalls()))
	assert.Equal(t, title, repository.RegisterCalls()[0].Title)
	assert.Equal(t, loginUserID, repository.RegisterCalls()[0].LoginUserID)
}

func Test_プロジェクトをDBへ登録(t *testing.T) {
	test.Init()
	r, err := app.NewProjectDbRepository()

	if err != nil {
		t.Fatal(err.Error())
	}

	title := "Test"
	id, _ := r.Register(title)

	assertデータが正しく登録されているか(t, id, title)
}

func assertデータが正しく登録されているか(t *testing.T, id int64, title string) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()

	if err != nil {
		t.Fatal(err)
	}

	row := db.QueryRow("SELECT * FROM projects WHERE id = $1", id)

	p := &app.ProjectRecord{}
	row.Scan(&p.ID, &p.Title, &p.CreatedAt, &p.UpdatedAt)

	assert.Equal(t, title, p.Title)
}

func Test_DBへの接続テスト(t *testing.T) {
	test.Init()
	r, _ := app.NewProjectDbRepository()

	result := r.Ping()

	assert.Equal(t, true, result)
}
