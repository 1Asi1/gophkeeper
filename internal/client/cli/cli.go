package cli

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gophkeeper/internal/client/models"
	"gophkeeper/internal/client/service"
)

type Cli struct {
	log   zerolog.Logger
	authS service.AuthService
	itemS service.ItemService
	root  *cobra.Command
}

func New(authS service.AuthService, itemS service.ItemService, log zerolog.Logger) *Cli {
	return &Cli{
		authS: authS,
		itemS: itemS,
		root: &cobra.Command{
			Use:   "gophkeeper",
			Short: "gophkeeper manager password",
			Long:  `gophkeeper password manager - graduate work for yandex workshop`,
		},
		log: log,
	}
}

func (c *Cli) Execute(ctx context.Context) error {
	var registerEmail string
	var registerPass string
	var register = &cobra.Command{
		Use:   "register",
		Short: "register user",
		Long:  `register new user`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.authS.Register(ctx, registerEmail, registerPass)
		},
	}
	register.Flags().StringVar(&registerEmail, "email", "", "user email")
	register.Flags().StringVar(&registerPass, "pass", "", "user password")

	var loginEmail string
	var loginPass string
	var login = &cobra.Command{
		Use:   "login",
		Short: "login",
		Long:  `login user`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.authS.Login(ctx, loginEmail, loginPass)
		},
	}
	login.Flags().StringVar(&loginEmail, "email", "", "user email")
	login.Flags().StringVar(&loginPass, "pass", "", "user password")

	var createItemType string
	var createItemData string
	var createItemMeta string
	var createItem = &cobra.Command{
		Use:   "create",
		Short: "create item",
		Long:  `create new item`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := c.itemS.Create(ctx, models.Item{
				Type: createItemType,
				Data: []byte(createItemData),
				Meta: []byte(createItemMeta),
			})
			if err != nil {
				return fmt.Errorf("c.itemS.Create: %w", err)
			}

			c.log.Info().Msgf("id item: %s", id.String())
			return nil
		},
	}
	createItem.Flags().StringVar(&createItemType, "type", "", "item type")
	createItem.Flags().StringVar(&createItemData, "data", "", "item data")
	createItem.Flags().StringVar(&createItemMeta, "meta", "", "item meta")

	var updateItemType string
	var updateItemData string
	var updateItemMeta string
	var updateItem = &cobra.Command{
		Use:   "update",
		Short: "update item",
		Long:  `update item`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.itemS.Update(ctx, models.Item{
				Type: updateItemType,
				Data: []byte(updateItemData),
				Meta: []byte(updateItemMeta),
			})
			if err != nil {
				return fmt.Errorf("c.itemS.Update: %w", err)
			}

			return nil
		},
	}
	updateItem.Flags().StringVar(&updateItemType, "type", "", "item type")
	updateItem.Flags().StringVar(&updateItemData, "data", "", "item data")
	updateItem.Flags().StringVar(&updateItemMeta, "meta", "", "item meta")

	var getAllItem = &cobra.Command{
		Use:   "all",
		Short: "get all",
		Long:  `get all items`,
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := c.itemS.GetAll(ctx, uuid.UUID{})
			if err != nil {
				return fmt.Errorf("c.itemS.GetAll: %w", err)
			}
			for _, v := range data {
				c.log.Info().Msgf("item: %+v\n", v)
			}

			return nil
		},
	}

	var getAllByType string
	var getAllByTypeItem = &cobra.Command{
		Use:   "all-type",
		Short: "get all by type",
		Long:  `get all by type items`,
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := c.itemS.GetAllByType(ctx, uuid.UUID{}, getAllByType)
			if err != nil {
				return fmt.Errorf("c.itemS.GetAllByType: %w", err)
			}
			for _, v := range data {
				c.log.Info().Msgf("item: %+v\n", v)
			}

			return nil
		},
	}
	getAllByTypeItem.Flags().StringVar(&getAllByType, "type", "", "item type")

	var getItemID string
	var getItem = &cobra.Command{
		Use:   "get",
		Short: "get",
		Long:  `get item`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := uuid.Parse(getItemID)
			if err != nil {
				return fmt.Errorf("uuid.Parse: %w", err)
			}

			data, err := c.itemS.Get(ctx, id)
			if err != nil {
				return fmt.Errorf("c.itemS.GetAllByType: %w", err)
			}

			c.log.Info().Msgf("item: %+v", data)
			return nil
		},
	}
	getItem.Flags().StringVar(&getItemID, "id", "", "item id")

	var deleteItemID string
	var deleteItem = &cobra.Command{
		Use:   "delete",
		Short: "delete",
		Long:  `delete item`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := uuid.Parse(getItemID)
			if err != nil {
				return fmt.Errorf("uuid.Parse: %w", err)
			}

			err = c.itemS.Delete(ctx, id)
			if err != nil {
				return fmt.Errorf("c.itemS.Delete: %w", err)
			}

			return nil
		},
	}
	deleteItem.Flags().StringVar(&deleteItemID, "id", "", "item id")

	c.root.AddCommand(register)
	c.root.AddCommand(login)
	c.root.AddCommand(createItem)
	c.root.AddCommand(updateItem)
	c.root.AddCommand(getAllItem)
	c.root.AddCommand(getAllByTypeItem)
	c.root.AddCommand(getItem)
	c.root.AddCommand(deleteItem)

	if err := c.root.Execute(); err != nil {
		return fmt.Errorf("rootCmd.Execute: %w", err)
	}

	return nil
}
