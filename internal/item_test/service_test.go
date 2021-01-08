package item_test

import (
	"github.com/mirodobrovocky/project1/internal/item"
	"github.com/mirodobrovocky/project1/internal/user"
	"reflect"
	"testing"
)

var items = []item.Item{
	item.NewItem("Item01", "User01", 199.99),
	item.NewItem("Item02", "User02", 299.99),
}

type fields struct {
	itemsRepository item.Repository
	userService 	user.Service
}

func Test_itemsService_Create(t *testing.T) {
	type args struct {
		create item.CreateDto
	}
	itemToBeCreated := item.NewItem("Item01", "me", 299.99)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *item.Item
		wantErr bool
	}{
		{
			name: "shouldCreateAndReturn",
			fields: fields{itemsRepositoryMock{}, userServiceMock{}},
			args: args{
				item.CreateDto{
					Name:  "Item01",
					Price: 299.99,
				},
			},
			want: &itemToBeCreated,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := item.NewService(tt.fields.itemsRepository, tt.fields.userService)
			got, err := s.Create(tt.args.create)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsService_FindAll(t *testing.T) {
	tests := []struct {
		name   	string
		fields 	fields
		want   	[]item.Item
		wantErr	bool
	}{
		{
			name: "shouldReturn",
			fields: fields{
				itemsRepositoryMock{}, userServiceMock{},
			},
			want: items,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := item.NewService(tt.fields.itemsRepository, tt.fields.userService)
			got, err := s.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsService_FindByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name   	string
		fields 	fields
		args   	args
		want   	*item.Item
		wantErr	bool
	}{
		{
			name: "shouldReturn",
			fields: fields{
				itemsRepositoryMock{}, userServiceMock{},
			},
			args: args{
				name: "foo",
			},
			want: &items[0],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := item.NewService(tt.fields.itemsRepository, tt.fields.userService)
			got, err := s.FindByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type itemsRepositoryMock struct {}

func (r itemsRepositoryMock) FindAll() ([]item.Item, error) {
	return items, nil
}

func (r itemsRepositoryMock) FindByName(name string) (*item.Item, error) {
	return &items[0], nil
}

func (r itemsRepositoryMock) Save(item item.Item) (*item.Item, error) {
	return &item, nil
}

type userServiceMock struct {

}

func (u userServiceMock) GetCurrentUser() (*user.User, error) {
	currentUser := user.New("me")
	return &currentUser, nil
}
