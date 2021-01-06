package service

import (
	"github.com/mirodobrovocky/project1/internal/item"
	"reflect"
	"testing"
)

var items = []item.Item{
	{
		Name: "Item01",
		Owner: "User01",
		Price: 199.99,
	},
	{
		Name: "Item02",
		Owner: "User02",
		Price: 299.99,
	},
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

func Test_itemsService_Create(t *testing.T) {
	type fields struct {
		itemsRepository item.Repository
	}
	type args struct {
		create item.CreateDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *item.Item
		wantErr bool
	}{
		{
			name: "shouldCreateAndReturn",
			fields: fields{itemsRepositoryMock{}},
			args: args{
				item.CreateDto{
					Name:  "Item01",
					Price: 299.99,
				},
			},
			want: &item.Item{
				Name:  "Item01",
				Owner: "CurrentUser",
				Price: 299.99,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := item.NewService(tt.fields.itemsRepository)
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
	type fields struct {
		itemsRepository item.Repository
	}
	tests := []struct {
		name   	string
		fields 	fields
		want   	[]item.Item
		wantErr	bool
	}{
		{
			name: "shouldReturn",
			fields: fields{
				itemsRepositoryMock{},
			},
			want: items,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := item.NewService(tt.fields.itemsRepository)
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
	type fields struct {
		itemsRepository item.Repository
	}
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
				itemsRepositoryMock{},
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
			s := item.NewService(tt.fields.itemsRepository)
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